import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class LogLevels {
  public static String message(String logLine) {
    String[] fields = logLine.split("\\s+", 2);
    if (fields.length > 1) {
      return fields[1].trim();
    } else {
      return "";
    }
  }

  public static String logLevel(String logLine) {
    Pattern r = Pattern.compile("\\[([a-zA-Z]+)\\].*");
    Matcher m = r.matcher(logLine);
    if (m.find()) {
      return m.group(1).toLowerCase();
    } else {
      return "";
    }
  }

  public static String reformat(String logLine) {
    return String.format("%s (%s)", message(logLine), logLevel(logLine));
  }
}
