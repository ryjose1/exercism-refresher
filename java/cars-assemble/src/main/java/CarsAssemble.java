import java.util.HashMap;

public class CarsAssemble {

  private static final HashMap<Integer, Integer> successRate = new HashMap<>();
  static {
      successRate.put(1, 100);
      successRate.put(2, 100);
      successRate.put(3, 100);
      successRate.put(4, 100);
      successRate.put(5, 90);
      successRate.put(6, 90);
      successRate.put(7, 90);
      successRate.put(8, 90);
      successRate.put(9, 80);
      successRate.put(10, 77);
  }

  private static final int DEFAULT_CARS_PER_HOUR = 221;

  public double productionRatePerHour(int speed) {
    if(successRate.containsKey(speed)) {
      return (double) DEFAULT_CARS_PER_HOUR * (double) successRate.get(speed) * (double) speed / 100;
    } else {
      return 0.0;
    }
  }

  public int workingItemsPerMinute(int speed) {
    return (int) (productionRatePerHour(speed) / 60);
  }
}
