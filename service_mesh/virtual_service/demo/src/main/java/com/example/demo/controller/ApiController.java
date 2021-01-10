package com.example.demo.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;

@RestController
public class ApiController {
    @GetMapping("/")
    public String getHome() {
        LocalDateTime dateTime = LocalDateTime.now();
        String version = System.getenv("VERSION");
        return "Currnet Date : " + dateTime.toString() + "\n version : " + version;
    }
}
