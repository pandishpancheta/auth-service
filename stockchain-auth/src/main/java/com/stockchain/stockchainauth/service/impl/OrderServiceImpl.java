package com.stockchain.stockchainauth.service.impl;

import com.stockchain.stockchainauth.dto.OrderDTO;
import com.stockchain.stockchainauth.service.OrderService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class OrderServiceImpl implements OrderService {
    @Override
    public List<OrderDTO> getAllOrders() {
        return null;
    }

    @Override
    public OrderDTO getOrderById(UUID id) {
        return null;
    }

    @Override
    public OrderDTO createOrder(OrderDTO orderDTO) {
        return null;
    }

    @Override
    public OrderDTO updateOrder(UUID id, OrderDTO orderDTO) {
        return null;
    }

    @Override
    public void deleteOrder(UUID id) {

    }
}
