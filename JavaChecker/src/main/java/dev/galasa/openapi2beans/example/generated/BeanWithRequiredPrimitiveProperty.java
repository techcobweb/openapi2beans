package dev.galasa.openapi2beans.example.generated;

// A bean with a single primitive property
public class BeanWithRequiredPrimitiveProperty {
    // Class Variables //
    private String aStringVariable;

    public BeanWithRequiredPrimitiveProperty (String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }

    // Getters //
    public String GetaStringVariable() {
        return aStringVariable;
    }

    // Setters //
    public void SetaStringVariable(String aStringVariable) {
        this.aStringVariable = aStringVariable;
    }
}