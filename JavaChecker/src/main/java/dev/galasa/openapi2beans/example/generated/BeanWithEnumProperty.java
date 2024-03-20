package dev.galasa.openapi2beans.example.generated;

// bean with a reference to an enum
public class BeanWithEnumProperty {
    // Class Variables //
    // an enum with 2 values to test against.
    private anEnumProperty anEnumProperty;

    public BeanWithEnumProperty (anEnumProperty anEnumProperty) {
        this.anEnumProperty = anEnumProperty;
    }

    // Getters //
    public anEnumProperty GetanEnumProperty() {
        return anEnumProperty;
    }

    // Setters //
    public void SetanEnumProperty(anEnumProperty anEnumProperty) {
        this.anEnumProperty = anEnumProperty;
    }
}