const byte writeReady = 17;
const byte readReady = 18;

bool readToWrite = false;
void setup() {
  //Setup serial
  Serial.begin(230400);
  Serial1.begin(230400);

  delay(1000);
}

void loop() {
  if (Serial.available() > 0) {
    byte message = Serial.read();
    if (message == writeReady) {
      Serial.write(writeReady);
      delay(100);
    } 
    else if (message == readReady) {

    }
    delay(1000);
    }
}

void write() {
  Serial.flush();
  Serial1.flush();
  for (int i = 0; i < 4000; i++) {
    if (Serial1.available() > 0) {
      byte data = Serial1.read();
      Serial.write(data);
    }
  }
}

