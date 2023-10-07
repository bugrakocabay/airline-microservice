package com.example.flightservice.repositories;

import com.example.flightservice.models.Flight;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.data.repository.query.Param;

import java.sql.Timestamp;
import java.util.Date;

public interface FlightRepository extends CrudRepository<Flight, Integer> {

    @Query("SELECT f FROM Flight f " +
            "WHERE (:arrival is null OR f.arrivalAirport = :arrival) " +
            "AND (:departure is null OR f.departureAirport = :departure) " +
            "AND (:from is null OR f.departure >= CAST(:from AS timestamp)) " +
            "AND (:to is null OR f.departure <= CAST(:to AS timestamp))")
    Iterable<Flight> findFlightsByOptionalParams(
            @Param("arrival") String arrival,
            @Param("departure") String departure,
            @Param("from") String from,
            @Param("to") String to);
}
