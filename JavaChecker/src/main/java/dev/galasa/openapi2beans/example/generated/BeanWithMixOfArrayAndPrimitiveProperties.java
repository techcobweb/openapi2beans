package dev.galasa.openapi2beans.example.generated;

// A bean with a mix of arrays and other primitive property types
public class BeanWithMixOfArrayAndPrimitiveProperties {
    // Constants //
    // Class Variables //
    private String[] anArrayVariable;
    private String aStringVariable;
    private int[] anIntArray;
    private int anIntVariable;

    public BeanWithMixOfArrayAndPrimitiveProperties () {
    }

    // Getters //
    public String[] GetAnArrayVariable() {
        return this.anArrayVariable;
    }
    public String GetAStringVariable() {
        return this.aStringVariable;
    }
    public int[] GetAnIntArray() {
        return this.anIntArray;
    }
    public int GetAnIntVariable() {
        return this.anIntVariable;
    }

    // Setters //
    public void SetAnArrayVariable(String[] anArrayVariable) {
        this.anArrayVariable = anArrayVariable;
    }
    public void SetAStringVariable(String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }
    public void SetAnIntArray(int[] anIntArray) {
        this.anIntArray = anIntArray;
    }
    public void SetAnIntVariable(int anIntVariable) {
        this.anIntVariable = anIntVariable;
    }
}