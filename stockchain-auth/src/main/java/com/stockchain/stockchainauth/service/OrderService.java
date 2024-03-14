package com.stockchain.stockchainauth.service;

import com.stockchain.stockchainauth.dto.OrderDTO;

import java.util.List;
import java.util.UUID;

public interface OrderService {
    List<OrderDTO> getAllOrders();
    OrderDTO getOrderById(UUID id);
    OrderDTO createOrder(OrderDTO orderDTO);
    OrderDTO updateOrder(UUID id, OrderDTO orderDTO);
    void deleteOrder(UUID id);
}
