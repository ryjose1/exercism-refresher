public class Lasagna {
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int elapsedMins) {
        return expectedMinutesInOven() - elapsedMins;
    }

    public int preparationTimeInMinutes(int layers) {
        return layers * 2;
    } 

    public int totalTimeInMinutes(int layers, int elapsedMins) {
        return preparationTimeInMinutes(layers) + elapsedMins;
    }
}
