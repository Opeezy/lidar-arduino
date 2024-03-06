
byte msp = 0;
byte lsp = 0;
byte shift = 1;
byte packet[2];
int count = 0;

void setup() {
  //Setup serial
  Serial.begin(9600);
}

void loop() {
  if (count != 11) {
      Serial.write(digitalRead(0));
      count++;
  } else {
    Serial.print("\n");
    count = 0;
  }
  delay(10);
}