package dev.galasa.openapi2beans.example.generated;

// A bean with a required property referencing another object
public class BeanWithRequiredReferencedObject {
    // Constants //
    // Class Variables //
    // An empty bean with no properties
    private EmptyBean referencingObject;

    public BeanWithRequiredReferencedObject () {
    }

    // Getters //
    public EmptyBean GetReferencingObject() {
        return this.referencingObject;
    }

    // Setters //
    public void SetReferencingObject(EmptyBean referencingObject) {
        this.referencingObject = referencingObject;
    }
}