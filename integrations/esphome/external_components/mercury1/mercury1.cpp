#include "mercury1.h"
#include "esphome/core/log.h"

namespace esphome {
  namespace mercury1 {
    static const char *const TAG = "mercury1";

    float Mercury1::get_setup_priority() const {
      return setup_priority::DATA;
    }

    void Mercury1::setup() {
      this->clean_uart_buffer();

      this->packet_generate(read_power_counters_request_, Command::READ_POWER_COUNTERS);
      this->packet_generate(read_params_current_request_, Command::READ_PARAMS_CURRENT);
    }

    void Mercury1::loop() {

    }

    void Mercury1::read_from_uart() {
      memset(this->read_buffer_, 0, MERCURY1_READ_BUFFER_SIZE);
      int response_len = 0;

      while (this->available()) {
        delay(10); // FIXME: задержка не портит буфер, без задержки байты читаются рандомно

        if(response_len > MERCURY1_READ_BUFFER_SIZE) {
          ESP_LOGW(TAG, "Buffer overflow");
          this->clean_uart_buffer();
          break;
        }

        this->read_byte(&this->read_buffer_[response_len]);
        response_len++;
      }

      ESP_LOGV(TAG, "Response raw %s", hexencode(this->read_buffer_, response_len).c_str());

      // игнорируем пакеты на отсылку команд самому счетчику
      if(response_len <= MERCURY1_READ_REQUEST_SIZE) {
        ESP_LOGD(TAG, "Skip response with length %d", response_len);
        return;
      }

      // обработка данных с валидных пакетов
      // TODO: на 206 счетчике с таймингами что-то не то и ответы склеиваются, поэтому вычитываем все

      for (uint8_t i = 0, begin = 0; i < response_len; begin = i) {
        memset(this->packet_buffer_, 0, MERCURY1_READ_BUFFER_SIZE);

        // --- process address ---
        if (i + MERCURY1_FIELD_ADDRESS_LENGTH > response_len) { // не достаточно длины для чтения адреса
          break;
        }

        memcpy(this->packet_buffer_, this->read_buffer_ + i, MERCURY1_FIELD_ADDRESS_LENGTH);

        // проверяем что адреса запроса и ответа совпадают
        if(memcmp(this->packet_buffer_, this->address_, MERCURY1_FIELD_ADDRESS_LENGTH) != 0) { // включительно, чтобы пропускать пакеты на отсылку команд
          ESP_LOGW(TAG, "Response first bytes isn't %s is %s", hexencode(this->address_, MERCURY1_FIELD_ADDRESS_LENGTH).c_str(), hexencode(this->packet_buffer_, MERCURY1_FIELD_ADDRESS_LENGTH).c_str());
          return;
        }

        // ADDR +4
        i += MERCURY1_FIELD_ADDRESS_LENGTH;

        // --- process command ---
        if (i + MERCURY1_FIELD_COMMAND_LENGTH > response_len) { // не достаточно длины для чтения команды
          break;
        }

        memcpy(this->packet_buffer_ + (i - begin), this->read_buffer_ + i, MERCURY1_FIELD_COMMAND_LENGTH);

        uint8_t cmd = this->read_buffer_[i];
        // CMD +1
        i += MERCURY1_FIELD_COMMAND_LENGTH;

        // --- process data ---
        uint8_t data_len = 0;
        uint8_t data_index = i;

        switch (cmd) {
            case Command::READ_POWER_COUNTERS:
              // T1(4)-T2(4)-T3(4)-T4(4)
              data_len = 4 * 4;
              break;

            case Command::READ_PARAMS_CURRENT:
              // V(2)-I(2)-P(3)
              data_len = 2 + 2 + 3;
              break;
        }

        if (data_len == 0) {
            ESP_LOGW(TAG, "Unknown response command 0x%02X", cmd);
            continue;
        }

        if (i + data_len > response_len) { // не достаточно длины для чтения данных
          break;
        }

        memcpy(this->packet_buffer_ + (i - begin), this->read_buffer_ + i, data_len);

        // DATA
        i += data_len;

        // --- process crc ---
        if (i + MERCURY1_FIELD_CRC_LENGTH > response_len) { // не достаточно длины для контрольных данных
          break;
        }

        memcpy(this->packet_buffer_ + (i - begin), this->read_buffer_ + i, MERCURY1_FIELD_CRC_LENGTH);

        // игнорируем пакеты с некорректной контрольной суммой, так как в эфире бывает дичь из обрывков пакетов
        uint16_t computed_crc = this->crc16(this->packet_buffer_, i - begin);
        uint16_t remote_crc = uint16_t(this->read_buffer_[i]) | (uint16_t(this->read_buffer_[i+1]) << 8);

        if (computed_crc != remote_crc) {
          ESP_LOGW(TAG, "CRC Check failed! computed %02X != remote %02X", computed_crc, remote_crc);
          return;
        }

        i += MERCURY1_FIELD_CRC_LENGTH;

        // --- debug ----
        ESP_LOGV(TAG, "Found valid packet %s", hexencode(this->packet_buffer_, i - begin).c_str());

        // --- update sensors ---
        switch (cmd) {
          case Command::READ_POWER_COUNTERS:
            this->T1 = this->to_long<4>(&this->read_buffer_[data_index]) * 10;
            this->T2 = this->to_long<4>(&this->read_buffer_[data_index+4]) * 10;
            this->T3 = this->to_long<4>(&this->read_buffer_[data_index+8]) * 10;
            this->T4 = this->to_long<4>(&this->read_buffer_[data_index+12]) * 10;
            this->TTotal = this->T1 +this->T2 + this->T3 + this->T4;

            this->tariff1_sensor_->publish_state(this->T1);
            this->tariff2_sensor_->publish_state(this->T2);
            this->tariff3_sensor_->publish_state(this->T3);
            this->tariff4_sensor_->publish_state(this->T4);
            this->tariffs_total_sensor_->publish_state(this->TTotal);
          break;

          case Command::READ_PARAMS_CURRENT:
            this->V = this->to_long(&this->read_buffer_[data_index]) / 10;
            this->A = this->to_long(&this->read_buffer_[data_index+2]) / 100.0;
            this->W = this->to_long<3>(&this->read_buffer_[data_index+4]);

            this->voltage_sensor_->publish_state(this->V);
            this->amperage_sensor_->publish_state(this->A);
            this->power_sensor_->publish_state(this->W);
          break;
        }
      }
    }

    void Mercury1::clean_uart_buffer() {
      while (this->available()) {
        this->read();
      }
    }

    void Mercury1::update() {
      ESP_LOGV(TAG, "Send READ_POWER_COUNTERS %s", hexencode(read_power_counters_request_, MERCURY1_READ_REQUEST_SIZE).c_str());
      this->write_array(read_power_counters_request_, MERCURY1_READ_REQUEST_SIZE);
      this->flush();

      delay(MERCURY1_WAIT_AFTER_SEND_REQUEST);

      this->read_from_uart();

      delay(MERCURY1_WAIT_AFTER_READ_RESPONSE);

      ESP_LOGV(TAG, "Send READ_PARAMS_CURRENT %s", hexencode(read_params_current_request_, MERCURY1_READ_REQUEST_SIZE).c_str());
      this->write_array(read_params_current_request_, MERCURY1_READ_REQUEST_SIZE);
      this->flush();

      delay(MERCURY1_WAIT_AFTER_SEND_REQUEST);

      this->read_from_uart();

      delay(MERCURY1_WAIT_AFTER_READ_RESPONSE);
    }

    void Mercury1::dump_config() {
      ESP_LOGCONFIG(TAG, "Mercury v1:");
      ESP_LOGCONFIG(TAG, "  Address %d", ((uint32_t) (this->address_[3] | (this->address_[2] << 8) | (this->address_[1] << 16) | (this->address_[0] << 24))));
      LOG_UPDATE_INTERVAL(this);
      LOG_SENSOR("", "Voltage", this->voltage_sensor_);
      LOG_SENSOR("", "Amperage", this->amperage_sensor_);
      LOG_SENSOR("", "Power", this->power_sensor_);
      LOG_SENSOR("", "Tariff 1", this->tariff1_sensor_);
      LOG_SENSOR("", "Tariff 2", this->tariff2_sensor_);
      LOG_SENSOR("", "Tariff 3", this->tariff3_sensor_);
      LOG_SENSOR("", "Tariff 4", this->tariff4_sensor_);
      LOG_SENSOR("", "Tariffs total", this->tariffs_total_sensor_);

      this->check_uart_settings(9600);
    }
  }  // namespace mercury1
}  // namespace esphome
