#include <SoftwareSerial.h>

const byte rxPin = 2;
const byte txPin = 3;

u_int8_t incomingByte = 0;

SoftwareSerial mySerial (rxPin, txPin);

void setup() {
  //Setup serial
  Serial.begin(230400);
  mySerial.begin(230400);
}

void loop() {
  incomingByte = mySerial.read();
  Serial.print(incomingByte, HEX);
}