package datastore

var RawJSON = `
[
  {
    "city": "Paris",
    "country": "France",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Tokyo",
    "country": "Japan",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "New York",
    "country": "USA",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "London",
    "country": "UK",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Berlin",
    "country": "Germany",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Rome",
    "country": "Italy",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Madrid",
    "country": "Spain",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Toronto",
    "country": "Canada",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Sydney",
    "country": "Australia",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Rio de Janeiro",
    "country": "Brazil",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Mumbai",
    "country": "India",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Beijing",
    "country": "China",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Moscow",
    "country": "Russia",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Seoul",
    "country": "South Korea",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Mexico City",
    "country": "Mexico",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Amsterdam",
    "country": "Netherlands",
    "clues": [
      "Known for its delicious local cuisine.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Stockholm",
    "country": "Sweden",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Zurich",
    "country": "Switzerland",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Cape Town",
    "country": "South Africa",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Buenos Aires",
    "country": "Argentina",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Istanbul",
    "country": "Turkey",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Bangkok",
    "country": "Thailand",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Cairo",
    "country": "Egypt",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Athens",
    "country": "Greece",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Dubai",
    "country": "UAE",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Lisbon",
    "country": "Portugal",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Brussels",
    "country": "Belgium",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Oslo",
    "country": "Norway",
    "clues": [
      "Known for its delicious local cuisine.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Copenhagen",
    "country": "Denmark",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Vienna",
    "country": "Austria",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Singapore",
    "country": "Singapore",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Kuala Lumpur",
    "country": "Malaysia",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Jakarta",
    "country": "Indonesia",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Hanoi",
    "country": "Vietnam",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Auckland",
    "country": "New Zealand",
    "clues": [
      "This city is famous for its historic landmarks.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Helsinki",
    "country": "Finland",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Dublin",
    "country": "Ireland",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Warsaw",
    "country": "Poland",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Prague",
    "country": "Czech Republic",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Budapest",
    "country": "Hungary",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Santiago",
    "country": "Chile",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Bogot\u00e1",
    "country": "Colombia",
    "clues": [
      "Known for its delicious local cuisine.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Lima",
    "country": "Peru",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Manila",
    "country": "Philippines",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Riyadh",
    "country": "Saudi Arabia",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Jerusalem",
    "country": "Israel",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Karachi",
    "country": "Pakistan",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Dhaka",
    "country": "Bangladesh",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Lagos",
    "country": "Nigeria",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Nairobi",
    "country": "Kenya",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Marrakech",
    "country": "Morocco",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Doha",
    "country": "Qatar",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Nur-Sultan",
    "country": "Kazakhstan",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Kyiv",
    "country": "Ukraine",
    "clues": [
      "Known for its delicious local cuisine.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Bucharest",
    "country": "Romania",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Belgrade",
    "country": "Serbia",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Bratislava",
    "country": "Slovakia",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Ljubljana",
    "country": "Slovenia",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Zagreb",
    "country": "Croatia",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Sofia",
    "country": "Bulgaria",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Vilnius",
    "country": "Lithuania",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Riga",
    "country": "Latvia",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Tallinn",
    "country": "Estonia",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Reykjavik",
    "country": "Iceland",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Luxembourg City",
    "country": "Luxembourg",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Valletta",
    "country": "Malta",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Nicosia",
    "country": "Cyprus",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Quito",
    "country": "Ecuador",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Montevideo",
    "country": "Uruguay",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Asunci\u00f3n",
    "country": "Paraguay",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "La Paz",
    "country": "Bolivia",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Caracas",
    "country": "Venezuela",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Colombo",
    "country": "Sri Lanka",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Yangon",
    "country": "Myanmar",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Kathmandu",
    "country": "Nepal",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Tashkent",
    "country": "Uzbekistan",
    "clues": [
      "Known for its delicious local cuisine.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Ashgabat",
    "country": "Turkmenistan",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Yerevan",
    "country": "Armenia",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Tbilisi",
    "country": "Georgia",
    "clues": [
      "This city is famous for its historic landmarks.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Beirut",
    "country": "Lebanon",
    "clues": [
      "This city is famous for its historic landmarks.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Muscat",
    "country": "Oman",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Manama",
    "country": "Bahrain",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Amman",
    "country": "Jordan",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Algiers",
    "country": "Algeria",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "More than 5 million people visit this city annually."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Tunis",
    "country": "Tunisia",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Addis Ababa",
    "country": "Ethiopia",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Accra",
    "country": "Ghana",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Abidjan",
    "country": "Ivory Coast",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Dakar",
    "country": "Senegal",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "One of the world's top universities is located here."
    ]
  },
  {
    "city": "Lusaka",
    "country": "Zambia",
    "clues": [
      "Home to one of the world's most famous rivers.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Gaborone",
    "country": "Botswana",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Antananarivo",
    "country": "Madagascar",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Havana",
    "country": "Cuba",
    "clues": [
      "This city hosts one of the biggest festivals in the world.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "Santo Domingo",
    "country": "Dominican Republic",
    "clues": [
      "Known for its delicious local cuisine.",
      "This city is famous for its historic landmarks."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "Tegucigalpa",
    "country": "Honduras",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "The city is home to a UNESCO World Heritage site.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  },
  {
    "city": "San Salvador",
    "country": "El Salvador",
    "clues": [
      "This city is famous for its historic landmarks.",
      "This city hosts one of the biggest festivals in the world."
    ],
    "fun_fact": [
      "The tallest building in the country is located here.",
      "This city has a unique underground tunnel system."
    ],
    "trivia": [
      "One of the world's top universities is located here.",
      "The city is home to a UNESCO World Heritage site."
    ]
  },
  {
    "city": "Panama City",
    "country": "Panama",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Known for its delicious local cuisine."
    ],
    "fun_fact": [
      "More than 5 million people visit this city annually.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "This city was once the capital of the country."
    ]
  },
  {
    "city": "San Jos\u00e9",
    "country": "Costa Rica",
    "clues": [
      "Known for its delicious local cuisine.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "The city has more bridges than any other in the world.",
      "It was originally a small settlement before growing into a major metropolis."
    ],
    "trivia": [
      "This city was once the capital of the country.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Guatemala City",
    "country": "Guatemala",
    "clues": [
      "This city is famous for its historic landmarks.",
      "A hub for culture, arts, and entertainment."
    ],
    "fun_fact": [
      "This city has a unique underground tunnel system.",
      "The city has more bridges than any other in the world."
    ],
    "trivia": [
      "The city's public transportation system is ranked among the top 10 globally.",
      "Its nickname is derived from an old legend."
    ]
  },
  {
    "city": "Belmopan",
    "country": "Belize",
    "clues": [
      "A hub for culture, arts, and entertainment.",
      "Home to one of the world's most famous rivers."
    ],
    "fun_fact": [
      "It was originally a small settlement before growing into a major metropolis.",
      "The tallest building in the country is located here."
    ],
    "trivia": [
      "Its nickname is derived from an old legend.",
      "The city's public transportation system is ranked among the top 10 globally."
    ]
  }
]`
