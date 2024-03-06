void setup() {
  //Setup serial
  Serial.begin(9600);
}

void loop() {
  Serial.write(digitalRead(0));
  delay(100);
}