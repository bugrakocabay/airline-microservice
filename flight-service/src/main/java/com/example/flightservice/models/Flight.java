package com.example.flightservice.models;

import jakarta.persistence.*;
import lombok.*;

import java.sql.Timestamp;

@Data
@Entity
@Builder
@Table(name = "flights")
@NoArgsConstructor
@AllArgsConstructor
public class Flight {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Integer flight_id;

    @Column(name = "flight_number")
    private String flightNumber;

    @Column(name = "departure_airport")
    private String departureAirport;

    @Column(name = "arrival_airport")
    private String arrivalAirport;

    @Column(name = "arrival")
    private Timestamp arrival;

    @Column(name = "departure")
    private Timestamp departure;

    @Column(name = "aircraft_id")
    private int aircraftId;

    @Column(name = "available_seats")
    private Integer availableSeats;

    private String status;
    private Integer capacity;
    private Integer price;
}
