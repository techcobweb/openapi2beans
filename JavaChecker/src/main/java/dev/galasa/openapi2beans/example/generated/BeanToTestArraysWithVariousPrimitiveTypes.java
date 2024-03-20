package dev.galasa.openapi2beans.example.generated;

// A bean to test array property can contain any primitive property type
public class BeanToTestArraysWithVariousPrimitiveTypes {
    // Class Variables //
    private double[] aNumberArray;
    private boolean[] aBooleanArray;
    private String[] aStringArray;
    private int[] anIntArray;

    public BeanToTestArraysWithVariousPrimitiveTypes () {
    }

    // Getters //
    public double[] GetaNumberArray() {
        return aNumberArray;
    }
    public boolean[] GetaBooleanArray() {
        return aBooleanArray;
    }
    public String[] GetaStringArray() {
        return aStringArray;
    }
    public int[] GetanIntArray() {
        return anIntArray;
    }

    // Setters //
    public void SetaNumberArray(double[] aNumberArray) {
        this.aNumberArray = aNumberArray;
    }
    public void SetaBooleanArray(boolean[] aBooleanArray) {
        this.aBooleanArray = aBooleanArray;
    }
    public void SetaStringArray(String[] aStringArray) {
        this.aStringArray = aStringArray;
    }
    public void SetanIntArray(int[] anIntArray) {
        this.anIntArray = anIntArray;
    }
}