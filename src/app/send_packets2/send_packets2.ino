const byte deviceOn = 2;
const byte appOn = 3;
const byte writeReady = 4;
const byte readReady = 5;

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
bool establishConnection() {
  
}
