package dev.galasa.openapi2beans.example.generated;

// A bean that tests arrays
public class BeanWithArrayOfReferenceToEmptyBean {
    // Class Variables //
    // An empty bean with no properties
    private EmptyBean[] anArrayVariable;

    public BeanWithArrayOfReferenceToEmptyBean () {
    }

    // Getters //
    public EmptyBean[] GetanArrayVariable() {
        return anArrayVariable;
    }

    // Setters //
    public void SetanArrayVariable(EmptyBean[] anArrayVariable) {
        this.anArrayVariable = anArrayVariable;
    }
}