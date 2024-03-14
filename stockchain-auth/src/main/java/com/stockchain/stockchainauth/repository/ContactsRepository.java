package com.stockchain.stockchainauth.repository;

import com.stockchain.stockchainauth.entity.Contacts;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface ContactsRepository extends JpaRepository<Contacts, UUID> { }
