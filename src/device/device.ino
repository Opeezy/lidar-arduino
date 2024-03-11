const byte deviceOn = 2;
const byte appOn = 3;
const byte writeReady = 4;
const byte readReady = 5;

bool connStatus = false;
void setup() {
  //Setup serial
  Serial.begin(230400);
  delay(1000);
}

void loop() {
  if (!connStatus) {
    establishConnection();
  } else {
    //TODO
  }
}
bool establishConnection() {
  //TODO: Send byte to app to establish connection
}
