#include <SoftwareSerial.h>

const u_int8_t rxPin = 2;
const u_int8_t txPin = 3;
const u_int8_t byteOne = 0x54;
const u_int8_t byteTwo = 0x2C;

SoftwareSerial mySerial (rxPin, txPin);

void setup() {
  //Setup serial
  Serial.begin(230400);
  mySerial.begin(230400);
  delay(1000);
}

void loop() {
  clearSerialBuffer();

  Serial.flush();-
  delay(50);
}

void clearSerialBuffer() {
  while (mySerial.available()) {
    byte next = mySerial.read();
    Serial.write(next);
  }
}
