package dev.galasa.openapi2beans.example.generated;

// A bean with a mix of arrays and other primitive property types
public class BeanWithMixOfArrayAndPrimitiveProperties {
    // Constants //
    // Class Variables //
    private int[] anIntArray;
    private int anIntVariable;
    private String[] anArrayVariable;
    private String aStringVariable;

    public BeanWithMixOfArrayAndPrimitiveProperties () {
    }

    // Getters //
    public int[] GetAnIntArray() {
        return this.anIntArray;
    }
    public int GetAnIntVariable() {
        return this.anIntVariable;
    }
    public String[] GetAnArrayVariable() {
        return this.anArrayVariable;
    }
    public String GetAStringVariable() {
        return this.aStringVariable;
    }

    // Setters //
    public void SetAnIntArray(int[] anIntArray) {
        this.anIntArray = anIntArray;
    }
    public void SetAnIntVariable(int anIntVariable) {
        this.anIntVariable = anIntVariable;
    }
    public void SetAnArrayVariable(String[] anArrayVariable) {
        this.anArrayVariable = anArrayVariable;
    }
    public void SetAStringVariable(String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }
}