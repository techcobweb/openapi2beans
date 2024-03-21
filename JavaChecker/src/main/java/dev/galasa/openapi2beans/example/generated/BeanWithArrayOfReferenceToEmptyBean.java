package dev.galasa.openapi2beans.example.generated;

// A bean that tests arrays
public class BeanWithArrayOfReferenceToEmptyBean {
    // Constants //
    // Class Variables //
    // An empty bean with no properties
    private EmptyBean[] anArrayVariable;

    public BeanWithArrayOfReferenceToEmptyBean () {
    }

    // Getters //
    public EmptyBean[] GetAnArrayVariable() {
        return this.anArrayVariable;
    }

    // Setters //
    public void SetAnArrayVariable(EmptyBean[] anArrayVariable) {
        this.anArrayVariable = anArrayVariable;
    }
}