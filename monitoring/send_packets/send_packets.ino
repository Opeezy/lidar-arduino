#include <SoftwareSerial.h>

const u_int8_t rxPin = 2;
const u_int8_t txPin = 3;
const u_int8_t byteOne = 0x54;
const u_int8_t byteTwo = 0x2C;

u_int8_t count = 0;

u_int8_t packet[47];

bool isPacket = false;
bool *ip = &isPacket;

SoftwareSerial mySerial (rxPin, txPin);

void setup() {
  //Setup serial
  Serial.begin(230400);
  mySerial.begin(230400);
  delay(1000);
}

void loop() {
  clearSerialBuffer();

  Serial.flush();
  mySerial.flush();

  delay(50);
}

void clearSerialBuffer() {
  while (mySerial.available()) {
    // Read next available byte
    u_int8_t outByte = mySerial.read();
    
    // First byte of packet should always be 84 DEC or 54 HEX
    if (outByte == byteOne && *ip == false) {
      // Set first byte of packet
      packet[0] = outByte;

      // Set to true to say we are writing a packet and increment
      *ip = true;
      count++;
    } else {
      
      // Check if we are currently building a packet
      if (*ip) {
        if (count == 46) {
          // Second byte of packet should always be 44 DEC or 0x2C HEX otherwise bad packet
          if (packet[1] == byteTwo) {
            // Fill last byte in packet
            packet[46] = outByte;

            // Write packet to serial
            writePacket();

            // Set isPacket to false and count to 0 to reset
            *ip = false;
            count = 0;
          }
          else {
            // Reset if bad packet
            count = 0;
            *ip = false;
          }
        } else {
          // Fill packet with current byte and increment count
          packet[count] = outByte;
          count++;
        }
      }
    }
  }
}

void writePacket() {
  // Loop through packet and write to Serial
  for (int i = 0; i < 47; i++) {
    // Print newline character if last byte else print
    Serial.write(packet[i]);
  }
}