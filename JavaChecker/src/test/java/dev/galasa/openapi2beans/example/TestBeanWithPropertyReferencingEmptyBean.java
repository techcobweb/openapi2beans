package dev.galasa.openapi2beans.example;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import dev.galasa.openapi2beans.example.generated.BeanWithPropertyReferencingEmptyBean;
import dev.galasa.openapi2beans.example.generated.EmptyBean;

public class TestBeanWithPropertyReferencingEmptyBean {
    
    @Test
    public void TestCanSerialiseTheBean() throws Exception {
        BeanWithPropertyReferencingEmptyBean beanUnderTest = new BeanWithPropertyReferencingEmptyBean();
        EmptyBean emptyBean = new EmptyBean();
        beanUnderTest.SetreferencingProperty(emptyBean);
        Gson gson = new GsonBuilder().setPrettyPrinting().create();
        String serialisedForm = gson.toJson(beanUnderTest);
        assertThat(serialisedForm).contains("\"referencingProperty\": {}");
    }
}
