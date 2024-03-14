package com.stockchain.stockchainauth.service.impl;

import com.stockchain.stockchainauth.dto.UserDTO;
import com.stockchain.stockchainauth.repository.UserRepository;
import com.stockchain.stockchainauth.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {
    private UserRepository userRepository;

    @Override
    public UserDTO getUserById(UUID id) {
        return null;
    }

    @Override
    public UserDTO getCurrentUser() {
        return null;
    }

    @Override
    public UserDTO updateUser(UUID id, UserDTO userDTO) {
        return null;
    }

    @Override
    public void deleteUser(UUID id) {

    }
}
