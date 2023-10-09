package com.example.flightservice.models;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Entity
@Builder
@Table(name = "aircrafts")
@NoArgsConstructor
@AllArgsConstructor
public class Aircraft {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "aircraft_id")
    private Integer aircraftId;

    @Column(name = "seating_capacity")
    private Integer seatingCapacity;

    @Column(name = "seating_configuration")
    private Integer seatingConfiguration;

    private String model;
}
