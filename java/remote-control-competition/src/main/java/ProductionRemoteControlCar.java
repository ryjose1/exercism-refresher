class ProductionRemoteControlCar implements RemoteControlCar, Comparable<ProductionRemoteControlCar> {

    private int distance;
    private int speed = 10;
    private int numVictories;

    public void drive() {
        this.distance += speed;
    }

    public int getDistanceTravelled() {
        return this.distance;
    }

    public int getNumberOfVictories() {
        return numVictories;
    }

    public void setNumberOfVictories(int numberofVictories) {
        this.numVictories = numberofVictories;
    }

    public int compareTo(ProductionRemoteControlCar car) {
        return this.getNumberOfVictories() - car.getNumberOfVictories();
    }
}
