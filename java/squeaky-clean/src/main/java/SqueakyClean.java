class SqueakyClean {
  private static final char ALPHA = '\u03B1';
  private static final char OMEGA = '\u03C9';

  static String clean(String identifier) {
    StringBuilder s = new StringBuilder();

    for (int i = 0; i < identifier.length(); i++) {
      char c = identifier.charAt(i);
      if(Character.isWhitespace(c)) {
        s.append('_');
      } else if (Character.isISOControl(c)) {
        s.append("CTRL");
      } else if (Character.isLetter(c) && (c < ALPHA || c > OMEGA)) {
        s.append(c);
      } else if (c == '-') {
        if (i+1 < identifier.length() && Character.isLetter(identifier.charAt(i+1))){
          s.append(Character.toUpperCase(identifier.charAt(++i)));
        }
      }
    }

    return s.toString();
  }
}
