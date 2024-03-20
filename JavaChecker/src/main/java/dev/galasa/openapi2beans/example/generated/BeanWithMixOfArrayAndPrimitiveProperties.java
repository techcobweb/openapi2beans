package dev.galasa.openapi2beans.example.generated;

// A bean with a mix of arrays and other primitive property types
public class BeanWithMixOfArrayAndPrimitiveProperties {
    // Class Variables //
    private int[] anIntArray;
    private int anIntVariable;
    private String[] anArrayVariable;
    private String aStringVariable;

    public BeanWithMixOfArrayAndPrimitiveProperties () {
    }

    // Getters //
    public int[] GetanIntArray() {
        return anIntArray;
    }
    public int GetanIntVariable() {
        return anIntVariable;
    }
    public String[] GetanArrayVariable() {
        return anArrayVariable;
    }
    public String GetaStringVariable() {
        return aStringVariable;
    }

    // Setters //
    public void SetanIntArray(int[] anIntArray) {
        this.anIntArray = anIntArray;
    }
    public void SetanIntVariable(int anIntVariable) {
        this.anIntVariable = anIntVariable;
    }
    public void SetanArrayVariable(String[] anArrayVariable) {
        this.anArrayVariable = anArrayVariable;
    }
    public void SetaStringVariable(String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }
}