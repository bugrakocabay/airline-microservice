package com.example.flightservice.dtos.requests;

import com.example.flightservice.models.Aircraft;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateFlightRequest {
    private String flightNumber;
    private String departureAirport;
    private String arrivalAirport;
    private Timestamp arrival;
    private Timestamp departure;
    private Aircraft aircraft;
    private Integer availableSeats;
    private String status;
    private Integer capacity;
    private Integer price;
    private Integer aircraftId;
}
