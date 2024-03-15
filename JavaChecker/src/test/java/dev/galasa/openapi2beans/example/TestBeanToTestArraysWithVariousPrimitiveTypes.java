package dev.galasa.openapi2beans.example;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import dev.galasa.openapi2beans.example.generated.BeanForTestingAllPrimitivePropertyTypes;

public class TestBeanToTestArraysWithVariousPrimitiveTypes {
    
    @Test
    public void TestCanSerialiseTheBean() throws Exception {
        BeanForTestingAllPrimitivePropertyTypes beanUnderTest = new BeanForTestingAllPrimitivePropertyTypes();
        beanUnderTest.SetaStringVariable("hello");
        beanUnderTest.SetaBooleanVariable(true);
        beanUnderTest.SetaIntVariable(11);
        beanUnderTest.SetaNumberVariable(1.28);
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(beanUnderTest);
        assertThat(serialisedForm).contains("\"aStringVariable\": \"hello\"");
        assertThat(serialisedForm).contains("\"aBooleanVariable\": true");
        assertThat(serialisedForm).contains("\"aIntVariable\": 11");
        assertThat(serialisedForm).contains("\"aNumberVariable\": 1.28");
    }
}
