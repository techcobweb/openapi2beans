package dev.galasa.openapi2beans.example;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import dev.galasa.openapi2beans.example.generated.BeanToTestArraysWithVariousPrimitiveTypes;

public class TestBeanToTestArraysWithVariousPrimitiveTypes {
    
    @Test
    public void TestCanSerialiseTheBean() throws Exception {
        BeanToTestArraysWithVariousPrimitiveTypes beanUnderTest = new BeanToTestArraysWithVariousPrimitiveTypes();
        beanUnderTest.SetAStringArray(new String[]{"randString0", "randString1"});
        beanUnderTest.SetABooleanArray(new boolean[]{true, false});
        beanUnderTest.SetAnIntArray(new int[]{2,3});
        beanUnderTest.SetANumberArray(new double[]{1.23, 4.56});
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(beanUnderTest);
        assertThat(serialisedForm).contains("\"aStringArray\": [\n" +
                        "    \"randString0\",\n" +
                        "    \"randString1\"\n" +
                        "  ]");
        assertThat(serialisedForm).contains("\"aBooleanArray\": [\n" +
                        "    true,\n" +
                        "    false\n" +
                        "  ]");
        assertThat(serialisedForm).contains("\"anIntArray\": [\n" +
                        "    2,\n" +
                        "    3\n" + 
                        "  ]");
        assertThat(serialisedForm).contains("\"aNumberArray\": [\n" +
                        "    1.23,\n" +
                        "    4.56\n" +
                        "  ]");
    }
}
