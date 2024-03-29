package dev.galasa.openapi2beans.example;


import org.junit.*;

import static org.assertj.core.api.Assertions.*;
import dev.galasa.openapi2beans.example.generated.*;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

public class TestBeanWithPrimitiveProperty {
    
    @Test
    public void TestCanSerialiseTheBean() throws Exception {
        BeanWithPrimitiveProperty beanUnderTest = new BeanWithPrimitiveProperty();
        beanUnderTest.SetAStringVariable("hello");
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(beanUnderTest);
        assertThat(serialisedForm).contains("\"aStringVariable\": \"hello\"");
    }
}