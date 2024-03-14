package com.stockchain.stockchainauth.service;

import com.stockchain.stockchainauth.dto.UserDTO;

import java.util.UUID;

public interface UserService {
    UserDTO getUserById(UUID id);
    UserDTO getCurrentUser();
    void deleteUser();
}
