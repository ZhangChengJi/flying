package com.example.demo;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoContoller {

//    @Autowired
//    private Grpc grpc;
@Value("${spring.datasource.type:}")
private String type;
    @GetMapping("/")
    public String get(){
     //   FlyingConfig flyingConfig=grpc.config("default");
       // grpc.shutdown();
      return  this.type;
    }
}

