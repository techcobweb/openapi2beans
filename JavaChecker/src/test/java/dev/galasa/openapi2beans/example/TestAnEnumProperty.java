package dev.galasa.openapi2beans.example;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import dev.galasa.openapi2beans.example.generated.anEnumProperty;


public class TestAnEnumProperty {
    
    
    @Test
    public void TestCanSerialiseTheEnumWithValue1() throws Exception {
        anEnumProperty enumUnderTest = anEnumProperty.string1;
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(enumUnderTest);
        assertThat(serialisedForm).contains("\"string1\"");
    }

    @Test
    public void TestCanSerialiseTheEnumWithValue2() throws Exception {
        anEnumProperty enumUnderTest = anEnumProperty.string2;
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(enumUnderTest);
        assertThat(serialisedForm).contains("\"string2\"");
    }
}
