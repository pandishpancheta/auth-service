package com.stockchain.stockchainauth.mapper;

import com.stockchain.stockchainauth.dto.UserDTO;
import com.stockchain.stockchainauth.entity.User;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.factory.Mappers;

@Mapper
public interface UserMapper {
    public UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    UserDTO toUserDTO(User user);
}
