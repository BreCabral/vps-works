package com.brenocabral.userauth.User;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
@RequestMapping("/users")
public class UserController {
    @GetMapping
    public String getAllUsers() {
        return "Hello World!";
    }
}
