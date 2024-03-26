package dev.galasa.openapi2beans.example.generated;

// A bean to test array property can contain any primitive property type
public class BeanToTestArraysWithVariousPrimitiveTypes {
    // Constants //
    // Class Variables //
    private boolean[] aBooleanArray;
    private String[] aStringArray;
    private int[] anIntArray;
    private double[] aNumberArray;

    public BeanToTestArraysWithVariousPrimitiveTypes () {
    }

    // Getters //
    public boolean[] GetABooleanArray() {
        return this.aBooleanArray;
    }
    public String[] GetAStringArray() {
        return this.aStringArray;
    }
    public int[] GetAnIntArray() {
        return this.anIntArray;
    }
    public double[] GetANumberArray() {
        return this.aNumberArray;
    }

    // Setters //
    public void SetABooleanArray(boolean[] aBooleanArray) {
        this.aBooleanArray = aBooleanArray;
    }
    public void SetAStringArray(String[] aStringArray) {
        this.aStringArray = aStringArray;
    }
    public void SetAnIntArray(int[] anIntArray) {
        this.anIntArray = anIntArray;
    }
    public void SetANumberArray(double[] aNumberArray) {
        this.aNumberArray = aNumberArray;
    }
}