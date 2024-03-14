package com.stockchain.stockchainauth.mapper;

import com.stockchain.stockchainauth.dto.ContactsDTO;
import com.stockchain.stockchainauth.entity.Contacts;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import java.util.List;

@Mapper
public interface ContactsMapper {
    ContactsMapper INSTANCE = Mappers.getMapper(ContactsMapper.class);

    ContactsDTO toContactsDTO(Contacts contacts);
    Contacts fromContactsDTO(ContactsDTO contactsDTO);
    List<ContactsDTO> toContactsDTOList(List<Contacts> contactsList);
    List<Contacts> fromContactsDTOList(List<ContactsDTO> contactsDTOList);
}
