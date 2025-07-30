package edu.northeastern.echo_service;

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/echo")
public class EchoServiceController {
    
    @PostMapping
    public String echo(@RequestBody String message) {
        return message;
    }
    
    @GetMapping
    public String echoGet(@RequestParam String message) {
        return message;
    }
}
