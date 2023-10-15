package com.example.flightservice.controllers;

import com.example.flightservice.dtos.requests.CreateFlightRequest;
import com.example.flightservice.models.Aircraft;
import com.example.flightservice.models.Flight;
import com.example.flightservice.repositories.AircraftRepository;
import com.example.flightservice.services.FlightService;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;

import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.verify;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@RunWith(SpringRunner.class)
@WebMvcTest(FlightController.class)
public class FlightControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private FlightService flightService;
    @MockBean
    private AircraftRepository aircraftRepository;

    @Test
    public void testGetAllFlights() throws Exception {
        // arrange
        List<Flight> mockFlightList = new ArrayList<>();
        mockFlightList.add(new Flight(
                1,
                "AA123",
                "CGK",
                "DPS",
                Timestamp.valueOf("2021-01-01 00:00:00"),
                Timestamp.valueOf("2021-01-02 00:00:00"),
                null,
                100,
                "ON TIME",
                100,
                500));
        mockFlightList.add(new Flight(
                2,
                "AA124",
                "CGK",
                "DPS",
                Timestamp.valueOf("2021-01-01 00:00:00"),
                Timestamp.valueOf("2021-01-02 00:00:00"),
                null,
                100,
                "ON TIME",
                100,
                500));

        given(flightService.getAllFlights()).willReturn(mockFlightList);

        // act
        mockMvc.perform(get("/flight")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error").value(false))
                .andExpect(jsonPath("$.message").value("OK"))
                .andExpect(jsonPath("$.data").isNotEmpty());
    }

    @Test
    public void testCreateFlight() throws Exception {
        // Arrange
        CreateFlightRequest createRequest = new CreateFlightRequest(
                "AA123",
                "CGK",
                "DPS",
                null,
                null,
                null,
                100,
                "ON TIME",
                100,
                500,
                1
        );
        Flight createdFlight = new Flight(
                1,
                "AA123",
                "CGK",
                "DPS",
                null,
                null,
                null,
                100,
                "ON TIME",
                100,
                500
        );
        given(flightService.createFlight(createRequest)).willReturn(createdFlight);
        given(aircraftRepository.findById(eq(1))).willReturn(Optional.of(new Aircraft(
                1,
                100,
                15,
                "Boeing 737")));

        // Act
        mockMvc.perform(post("/flight")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content("""
                                {
                                    "flightNumber": "AA123",
                                    "departureAirport": "CGK",
                                    "arrivalAirport": "DPS",
                                    "arrival": "2021-01-01T00:00:00",
                                    "departure": "2021-01-02T00:00:00",
                                    "aircraft": null,
                                    "availableSeats": 100,
                                    "status": "ON TIME",
                                    "capacity": 100,
                                    "price": 500,
                                    "aircraftId": 1
                                }"""))
                .andExpect(status().isCreated())
                .andExpect(jsonPath("$.error").value(false))
                .andExpect(jsonPath("$.message").value("OK"))
                .andExpect(jsonPath("$.data.flightId").value(createdFlight.getFlightId()))
                .andExpect(jsonPath("$.data.flightNumber").value(createdFlight.getFlightNumber()));

        verify(flightService).createFlight(createRequest);
    }
}
