class NeedForSpeed {
    private int battery = 100;
    private int distance;
    private int speed;
    private int batteryDrain;

    public NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public int getSpeed() {
        return speed;
    }

    public int getBatteryDrain() {
        return batteryDrain;
    }

    public int getBattery() {
        return battery;
    }

    public boolean batteryDrained() {
        return this.battery == 0;

    }

    public int distanceDriven() {
        return this.distance;
    }

    public void drive() {
        if (this.batteryDrain > this.battery) {
            this.battery = 0;
        } else {
            this.distance += this.speed;
            this.battery -= this.batteryDrain;
        }
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {
    private int distance;
    public RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean carCanFinish(NeedForSpeed car) {
        return this.distance <= car.getBattery() / car.getBatteryDrain() * car.getSpeed();
    }
}
