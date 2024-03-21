package dev.galasa.openapi2beans.example;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import dev.galasa.openapi2beans.example.generated.BeanWithEnumProperty;
import dev.galasa.openapi2beans.example.generated.AnEnumProperty;


public class TestBeanWithEnumProperty {
    
    @Test
    public void TestCanSerialiseTheBean() throws Exception {
        AnEnumProperty enumProperty = AnEnumProperty.string1;
        BeanWithEnumProperty beanUnderTest = new BeanWithEnumProperty(enumProperty);
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(beanUnderTest);
        assertThat(serialisedForm).contains("\"anEnumProperty\": \"string1\"");
    }
}
