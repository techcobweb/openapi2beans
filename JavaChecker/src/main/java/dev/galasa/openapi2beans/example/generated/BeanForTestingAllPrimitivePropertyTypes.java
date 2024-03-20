package dev.galasa.openapi2beans.example.generated;

// A bean that tests all primitive property types (except arrays)
public class BeanForTestingAllPrimitivePropertyTypes {
    // Class Variables //
    // this should be a String in a java bean
    private String aStringVariable;
    // this should be an int in a java bean
    private int aIntVariable;
    // this should be a float in the java bean
    private double aNumberVariable;
    private boolean aBooleanVariable;

    public BeanForTestingAllPrimitivePropertyTypes () {
    }

    // Getters //
    public String GetaStringVariable() {
        return aStringVariable;
    }
    public int GetaIntVariable() {
        return aIntVariable;
    }
    public double GetaNumberVariable() {
        return aNumberVariable;
    }
    public boolean GetaBooleanVariable() {
        return aBooleanVariable;
    }

    // Setters //
    public void SetaStringVariable(String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }
    public void SetaIntVariable(int aIntVariable) {
        this.aIntVariable = aIntVariable;
    }
    public void SetaNumberVariable(double aNumberVariable) {
        this.aNumberVariable = aNumberVariable;
    }
    public void SetaBooleanVariable(boolean aBooleanVariable) {
        this.aBooleanVariable = aBooleanVariable;
    }
}