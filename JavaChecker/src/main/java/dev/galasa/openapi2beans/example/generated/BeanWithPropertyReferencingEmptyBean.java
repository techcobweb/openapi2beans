package dev.galasa.openapi2beans.example.generated;

public class BeanWithPropertyReferencingEmptyBean {
    // Constants //
    // Class Variables //
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