package com.example.flightservice.services;

import com.example.flightservice.models.Flight;
import com.example.flightservice.repositories.FlightRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.sql.Date;
import java.text.ParseException;
import java.text.SimpleDateFormat;

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
            java.util.Date from,
            java.util.Date to
    ) throws ParseException {
        java.sql.Date fromDateSql = null;
        java.sql.Date toDateSql = null;

        // Convert "from" string and "to" string to Date objects
        SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS");
        if (from != null) {
            // java.util.Date fromDateUtil = dateFormat.parse(from);
            fromDateSql = new java.sql.Date(from.getTime());
        }
        if (to != null) {
            // java.util.Date toDateUtil = dateFormat.parse(to);
            toDateSql = new java.sql.Date(to.getTime());
        }

        return flightRepository.findFlightsByOptionalParams(arrival, departure, fromDateSql, toDateSql);
    }
}
