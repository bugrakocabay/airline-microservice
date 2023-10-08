package com.example.flightservice.services;

import com.example.flightservice.dtos.responses.BaseResponse;
import com.example.flightservice.models.Flight;
import com.example.flightservice.repositories.FlightRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Optional;

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

    @Transactional
    public Flight getFlightById(Integer id) {
        return flightRepository.findById(id).orElse(null);
    }

    @Transactional
    public Flight updateFlightById(Integer id, Flight flight) {
        return flightRepository.findById(id)
                .map(existingFlight -> {
                    Optional.ofNullable(flight.getArrivalAirport()).ifPresent(existingFlight::setArrivalAirport);
                    Optional.ofNullable(flight.getDepartureAirport()).ifPresent(existingFlight::setDepartureAirport);
                    Optional.ofNullable(flight.getDeparture()).ifPresent(existingFlight::setDeparture);
                    Optional.ofNullable(flight.getArrival()).ifPresent(existingFlight::setArrival);
                    Optional.ofNullable(flight.getAvailableSeats()).ifPresent(existingFlight::setAvailableSeats);
                    Optional.ofNullable(flight.getPrice()).ifPresent(existingFlight::setPrice);
                    Optional.ofNullable(flight.getCapacity()).ifPresent(existingFlight::setCapacity);
                    Optional.ofNullable(flight.getStatus()).ifPresent(existingFlight::setStatus);

                    return flightRepository.save(existingFlight);
                })
                .orElse(null);
    }

    @Transactional
    public void deleteFlightById(Integer id) {
        flightRepository.deleteById(id);
    }
}
