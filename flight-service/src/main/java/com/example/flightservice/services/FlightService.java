package com.example.flightservice.services;

import com.example.flightservice.dtos.requests.CreateFlightRequest;
import com.example.flightservice.models.Aircraft;
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
    @Autowired
    private AircraftService aircraftService;

    @Transactional
    public Iterable<Flight> getAllFlights() {
        return flightRepository.findAll();
    }

    @Transactional
    public Flight createFlight(CreateFlightRequest flight)  {
        Aircraft aircraft = aircraftService.getAircraftById(flight.getAircraftId());
        if (aircraft == null) {
            return null;
        }

        return flightRepository.save(
                Flight.builder()
                        .flightNumber(flight.getFlightNumber())
                        .departureAirport(flight.getDepartureAirport())
                        .arrivalAirport(flight.getArrivalAirport())
                        .arrival(flight.getArrival())
                        .departure(flight.getDeparture())
                        .aircraft(aircraft)
                        .availableSeats(flight.getAvailableSeats())
                        .status(flight.getStatus())
                        .capacity(flight.getCapacity())
                        .price(flight.getPrice())
                        .build()
        );
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
