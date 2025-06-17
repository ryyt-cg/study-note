API Key
Key - 27ADI5HR3OMBHWEN
Secret - KGoy6cxg7+SSVbEk2cJ5LYFBL8LZN6z+AtfDDgz4b4vZTFW4jskDaevzPqo/71Oo

+------------+------------------------------------------------------------------+
| API Key | 2NKOCIXYZQGCKRJR |
| API Secret | cer/09AhNqD9wVsm1ykb+H/TUWFTY0qZ84ghzqeAPLcHhIB9potpY/MQmP/4crG8 |
+------------+------------------------------------------------------------------+

code snippet

```bash
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic MjdBREk1SFIzT01CSFdFTjpLR295NmN4ZzcrU1NWYkVrMmNKNUxZRkJMOExaTjZ6K0F0ZkREZ3o0YjR2WlRGVzRqc2tEYWV2elBxby83MU9v" \
  https://pkc-n98pk.us-west-2.aws.confluent.cloud:443/kafka/v3/clusters/lkc-vy3zrz/topics \
  -d '{"topic_name":"weblog"}'
```

application.properties

```properties
# Required connection configs for Kafka producer, consumer, and admin
spring.kafka.properties.sasl.mechanism=PLAIN
spring.kafka.bootstrap-servers=pkc-n98pk.us-west-2.aws.confluent.cloud:9092
spring.kafka.properties.sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username='27ADI5HR3OMBHWEN' password='KGoy6cxg7+SSVbEk2cJ5LYFBL8LZN6z+AtfDDgz4b4vZTFW4jskDaevzPqo/71Oo';
spring.kafka.properties.security.protocol=SASL_SSL

# Best practice for higher availability in Apache Kafka clients prior to 3.0
spring.kafka.properties.session.timeout.ms=45000
```

build.gradle

```gradle
plugins {
  id 'org.springframework.boot' version "2.7.5"
  id 'io.spring.dependency-management' version '1.1.0'
  id 'java'
}

repositories {
  mavenCentral()
}

group = 'example'
version = '0.0.1-SNAPSHOT'

java {
  sourceCompatibility = '17'
}

dependencies {
  implementation 'org.springframework.boot:spring-boot-starter-web:2.7.5'
  implementation 'org.apache.kafka:kafka-clients'
  implementation 'org.springframework.kafka:spring-kafka'
}
```

```java
package example;

import example.Producer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.listener.MessageListenerContainer;
import org.springframework.kafka.config.KafkaListenerEndpointRegistry;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.WebApplicationType;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class SpringBootWithKafkaApplication {

  private final Producer producer;

  public static void main(String[] args) {
    SpringApplication application = new SpringApplication(SpringBootWithKafkaApplication.class);
    application.setWebApplicationType(WebApplicationType.NONE);
    application.run(args);
  }

  @Bean
  public CommandLineRunner CommandLineRunnerBean() {
    return (args) -> {
      this.producer.sendMessage("key", "value");
      MessageListenerContainer listenerContainer = kafkaListenerEndpointRegistry.getListenerContainer("myConsumer");
      listenerContainer.start();
    };
  }

  @Autowired
  SpringBootWithKafkaApplication(Producer producer) {
    this.producer = producer;
  }

  @Autowired
  private KafkaListenerEndpointRegistry kafkaListenerEndpointRegistry;

}


package example;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;
import org.springframework.util.concurrent.ListenableFuture;
import org.springframework.util.concurrent.ListenableFutureCallback;
import org.springframework.kafka.support.SendResult;

@Service
public class Producer {

  private static final Logger logger = LoggerFactory.getLogger(Producer.class);
  private static final String TOPIC = "weblog";

  @Autowired
  private KafkaTemplate<String, String> kafkaTemplate;

  public void sendMessage(String key, String value) {
    ListenableFuture<SendResult<String, String>> future = kafkaTemplate.send(TOPIC, key, value);
    future.addCallback(new ListenableFutureCallback<SendResult<String, String>>() {
      @Override
      public void onSuccess(SendResult<String, String> result) {
        logger.info(String.format("\n\n Produced event to topic %s: key = %-10s value = %s \n\n", TOPIC, key, value));
      }

      @Override
      public void onFailure(Throwable ex) {
        ex.printStackTrace();
      }
    });
  }
}



package example;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;

import java.io.IOException;

@Service
public class Consumer {

  private final Logger logger = LoggerFactory.getLogger(Consumer.class);

  @KafkaListener(id = "myConsumer", topics = "weblog", groupId = "springboot-group-1", autoStartup = "false")
  public void listen(String value,
    @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
    @Header(KafkaHeaders.RECEIVED_MESSAGE_KEY) String key) {
    logger.info(String.format("\n\n Consumed event from topic %s: key = %-10s value = %s \n\n", topic, key, value));
  }
}
```


