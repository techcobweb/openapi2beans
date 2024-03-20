package dev.galasa.openapi2beans.example.generated;

// A bean to test referencing functionality
public class BeanWithPropertyReferencingEmptyBean {
    // Class Variables //
    // An empty bean with no properties
    private EmptyBean referencingProperty;

    public BeanWithPropertyReferencingEmptyBean () {
    }

    // Getters //
    public EmptyBean GetreferencingProperty() {
        return referencingProperty;
    }

    // Setters //
    public void SetreferencingProperty(EmptyBean referencingProperty) {
        this.referencingProperty = referencingProperty;
    }
}