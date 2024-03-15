void setup() {
  //Setup serial
  Serial.begin(230400);
  Serial1.begin(230400);

  delay(1000);
}

void loop() {
  while (Serial1.available()) {
    byte data = Serial1.read();
    Serial.write(data);
  }
}


