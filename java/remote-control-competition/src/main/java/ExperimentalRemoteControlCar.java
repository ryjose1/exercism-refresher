public class ExperimentalRemoteControlCar implements RemoteControlCar{

    private int distance;
    private int speed = 20;

    public void drive() {
        this.distance += this.speed;
    }

    public int getDistanceTravelled() {
        return this.distance;
    }
}
