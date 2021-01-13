package com.example.front.controller;

import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@RestController
public class FrontController {
    private final static String URL = "http://webservice:80/";

    @GetMapping("/hello")
    public ResponseEntity<String> callRemoteService() {
        HttpHeaders header = new HttpHeaders();
        HttpEntity<?> entity = new HttpEntity<>(header);

        RestTemplate resttemplate = new RestTemplate();
        return resttemplate.getForEntity(URL, String.class);
    }
}
