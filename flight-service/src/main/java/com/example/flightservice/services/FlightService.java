package com.example.flightservice.services;

import com.example.flightservice.models.Flight;
import com.example.flightservice.repositories.FlightRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.sql.Timestamp;
import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
import java.util.Date;

@Service
public class FlightService {

    @Autowired
    private FlightRepository flightRepository;

    @Transactional
    public Iterable<Flight> getAllFlights() {
        return flightRepository.findAll();
    }

    @Transactional
    public Flight createFlight(Flight flight) {
        return flightRepository.save(flight);
    }

    @Transactional
    public Iterable<Flight> searchFlight(
            String arrival,
            String departure,
            Date from,
            Date to) {
        String fromString = null;
        String toString = null;

        DateFormat formatter = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        if (from != null) {
            fromString = formatter.format(from);
        }
        if (to != null) {
            toString = formatter.format(to);
        }

        return flightRepository.findFlightsByOptionalParams(arrival, departure, fromString, toString);
    }
}
